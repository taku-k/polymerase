package cli

import (
	"net"

	"github.com/spf13/cobra"

	"github.com/taku-k/polymerase/pkg/base"
	"github.com/taku-k/polymerase/pkg/polypb"
	"github.com/taku-k/polymerase/pkg/utils"
)

var serverConnHost, serverConnPort, serverAdvertiseHost string
var clientConnHost, clientConnPort string

var serverCfg = base.MakeServerConfig()
var baseCfg = serverCfg.Config
var backupCtx = backupContext{Config: baseCfg}
var restoreCtx = MakeRestoreContext(baseCfg)
var xtrabackupCfg = base.MakeXtrabackupConfig()

func initXtrabackupConfig() {
	xtrabackupCfg.XtrabackupBinPath =
		utils.EnvOrDefaultString("POLYMERASE_XTRABACKUP_PATH", xtrabackupCfg.XtrabackupBinPath)
}

func init() {
	startCmd.PersistentPreRun = func(cmd *cobra.Command, _ []string) {
		baseCfg.Host = serverConnHost
		baseCfg.Port = serverConnPort
		baseCfg.Addr = net.JoinHostPort(serverConnHost, serverConnPort)
		baseCfg.AdvertiseAddr = net.JoinHostPort(serverAdvertiseHost, serverConnPort)
	}

	fullBackupCmd.PersistentPreRunE = func(cmd *cobra.Command, _ []string) error {
		baseCfg.Addr = net.JoinHostPort(clientConnHost, clientConnPort)
		backupCtx.backupType = polypb.BackupType_XTRABACKUP_FULL
		initXtrabackupConfig()
		return nil
	}

	incBackupCmd.PersistentPreRunE = func(cmd *cobra.Command, _ []string) error {
		baseCfg.Addr = net.JoinHostPort(clientConnHost, clientConnPort)
		backupCtx.backupType = polypb.BackupType_XTRABACKUP_INC
		initXtrabackupConfig()
		return nil
	}

	mysqldumpCmd.PersistentPreRunE = func(cmd *cobra.Command, _ []string) error {
		baseCfg.Addr = net.JoinHostPort(clientConnHost, clientConnPort)
		backupCtx.backupType = polypb.BackupType_MYSQLDUMP
		return nil
	}

	restoreCmd.PersistentPreRunE = func(cmd *cobra.Command, _ []string) error {
		baseCfg.Addr = net.JoinHostPort(clientConnHost, clientConnPort)
		initXtrabackupConfig()
		xtrabackupCfg.UseMemory = restoreCtx.useMemory.String()
		return nil
	}

	// Client Flags
	clientCmds := []*cobra.Command{
		fullBackupCmd,
		incBackupCmd,
		mysqldumpCmd,
		restoreCmd,
	}

	for _, cmd := range clientCmds {
		f := cmd.Flags()

		f.StringVar(&clientConnHost, "host", "127.0.0.1", "Polymerase server hostname.")
		f.StringVar(&clientConnPort, "port", "24925", "Polymerase server port.")
	}

	// Backup and restore commands flags
	for _, cmd := range []*cobra.Command{
		fullBackupCmd,
		incBackupCmd,
		restoreCmd,
	} {
		f := cmd.Flags()

		f.VarP(&backupCtx.db, "db", "d", "Database ID")
		f.StringVar(&xtrabackupCfg.DefaultsFile, "defaults-file", xtrabackupCfg.DefaultsFile, "Read default MySQL options from the given file.")
	}

	// Backup commands flags
	for _, cmd := range []*cobra.Command{fullBackupCmd, incBackupCmd} {
		f := cmd.PersistentFlags()

		f.StringVar(&xtrabackupCfg.Host, "mysql-host", xtrabackupCfg.Host, "The MySQL hostname to connect with.")
		f.StringVarP(&xtrabackupCfg.Port, "mysql-port", "p", xtrabackupCfg.Port, "The MySQL port to connect with.")
		f.StringVarP(&xtrabackupCfg.User, "mysql-user", "u", xtrabackupCfg.User, "The MySQL username to connect with.")
		f.StringVarP(&xtrabackupCfg.Password, "mysql-password", "P", xtrabackupCfg.Password, "The MySQL password to connect with.")
		f.BoolVar(&xtrabackupCfg.InsecureAuth, "insecure-auth", xtrabackupCfg.InsecureAuth, "Connect with insecure auth. It is useful when server uses old protocol.")
		f.IntVar(&xtrabackupCfg.Parallel, "parallel", xtrabackupCfg.Parallel, "The number of threads to use to copy multiple data files concurrently when creating a backup.")
		f.StringVar(&backupCtx.compressCmd, "compress-cmd", "gzip -c", "Use external compression program command.")
	}

	{
		f := mysqldumpCmd.PersistentFlags()

		f.VarP(&backupCtx.db, "db", "d", "Database ID")
		f.StringVar(&xtrabackupCfg.Host, "mysql-host", xtrabackupCfg.Host, "The MySQL hostname to connect with.")
		f.StringVarP(&xtrabackupCfg.Port, "mysql-port", "p", xtrabackupCfg.Port, "The MySQL port to connect with.")
		f.StringVarP(&xtrabackupCfg.User, "mysql-user", "u", xtrabackupCfg.User, "The MySQL username to connect with.")
		f.StringVarP(&xtrabackupCfg.Password, "mysql-password", "P", xtrabackupCfg.Password, "The MySQL password to connect with.")
	}

	// Full-backup command specific
	{
		f := fullBackupCmd.Flags()

		f.BoolVar(&backupCtx.purgePrev, "purge-prev", false, "The flag whether previous backups are purged.")
	}

	// Restore command specific
	{
		f := restoreCmd.Flags()

		f.StringVar(&restoreCtx.fromStr, "fromStr", restoreCtx.fromStr, "")
		f.BoolVar(&restoreCtx.applyPrepare, "apply-prepare", restoreCtx.applyPrepare, "")
		f.Var(&restoreCtx.maxBandWidth, "max-bandwidth", "max bandwidth for download src archives (Bytes/sec)")
		f.BoolVar(&restoreCtx.latest, "latest", false, "Fetch the latest backups.")
		f.StringVar(&restoreCtx.decompressCmd, "decompress-cmd", "gzip", "Use external decompression program command")
		f.Var(&restoreCtx.useMemory, "use-memory", "How much memory is allocated for preparing a backup.")
	}

	// Start Flags
	{
		f := startCmd.Flags()

		f.StringVar(&serverConnHost, "host", string(serverCfg.HostName), "The hostname to listen on.")
		f.StringVar(&serverAdvertiseHost, "advertise-host", string(serverCfg.HostName), "The hostname to advertise to other nodes and clients.")
		f.StringVar(&serverConnPort, "port", base.DefaultPort, "The port to bind to.")
		f.Var(serverCfg.StoreDir, "store-dir", "The dir path to store data files.")
		f.StringVar(&serverCfg.JoinAddr, "join", "", "The address of node which acts as bootstrap when joining an existing cluster.")
		f.StringVar(&serverCfg.EtcdPeerPort, "etcd-peer-port", "2380", "The port to be used for etcd peer communication.")
		f.Var(&serverCfg.NodeID, "name", "The human-readable name.")
	}

	rootCmd.AddCommand(
		startCmd,
		fullBackupCmd,
		incBackupCmd,
		mysqldumpCmd,
		restoreCmd,
		versionCmd)
}
