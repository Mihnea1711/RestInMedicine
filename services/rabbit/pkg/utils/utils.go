package utils

func GetDirection(consumerFlag bool) string {
	if consumerFlag {
		return QueueDirectionListen
	} else {
		return QueueDirectionPublish
	}
}
