package sync

func (connection *Connection) CommandBuilder(command string, args ...string) []interface{} {
	var ret []interface{}

	switch connection.GetType() {
	case "local":
		ret = connection.LocalCommandBuilder(command, args...)
	case "ssh":
		ret = connection.SshCommandBuilder(command, args...)
	case "ssh+docker":
		fallthrough
	case "docker":
		ret = connection.DockerCommandBuilder(command, args...)
	default:
		panic(connection)
	}

	return ret
}

func (connection *Connection) GetType() string {
	var connType string

	// autodetection
	if (connection.Type == "") || (connection.Type == "auto") {
		connection.Type = "local"

		if (connection.Docker != "") && connection.Hostname != "" {
			connection.Type = "ssh+docker"
		} else if connection.Docker != "" {
			connection.Type = "docker"
		} else if connection.Hostname != "" {
			connection.Type = "ssh"
		}
	}

	switch connection.Type {
	case "local":
		connType = "local"
	case "ssh":
		connType = "ssh"
	case "docker":
		connType = "docker"
	case "ssh+docker":
		connType = "ssh+docker"
	default:
		Logger.FatalExit(1, "Unknown connection type \"%s\"", connType)
	}

	return connType
}
