package utils

const CONFIG_PATH = "configs/config.yaml"
const CLEAR_DB_RESOURCES_TIMEOUT = 10
const RABBIT_CLOSE_TIMEOUT = 5

const (
	// QueueDirectionListen represents a queue for listening/consuming
	QueueDirectionListen string = "listen"
	// QueueDirectionPublish represents a queue for publishing
	QueueDirectionPublish string = "publish"
)

const (
	DELETE_QUEUE = "delete_queue"
)
