package core_constant

const (
	NAME_SHORT   = "BKT"
	NAME         = "BakaTools"
	PACKAGE_NAME = "bakatools"
	CLI_COMMAND  = PACKAGE_NAME

	VERSION = "1.0.0" // semver (MAJOR.MINOR.PATCH)

	GITHUB_URL        = "https://github.com/Simp-Community/bakatools"
	GITHUB_API_LATEST = "https://api.github.com/repos/Simp-Community/bakatools/releases/latest"
	DOCUMENTATION_URL = "https://bakatools.readthedocs.io"

	PACKAGE_PATH         = ""
	LOGGING_FILE         = "logs/bakatools.log"
	LANGUAGE_FILE_SUFFIX = ".yml"
	DEFAULT_LANGUAGE     = "en_us"

	PLUGIN_THREAD_POOL_SIZE               = 4
	MAX_TASK_QUEUE_SIZE                   = 2048
	WAIT_TIME_AFTER_SERVER_STDOUT_END_SEC = 60
	REACTOR_QUEUE_FULL_WARN_INTERVAL_SEC  = 5
)
