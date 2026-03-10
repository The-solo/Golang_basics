package main

type notification interface {
	importance() int
}

type directMessage struct {
	senderUsername string
	messageContent string
	priorityLevel  int
	isUrgent       bool
}

type groupMessage struct {
	groupName      string
	messageContent string
	priorityLevel  int
}

type systemAlert struct {
	alertCode      string
	messageContent string
}

func (d directMessage) importance() int{
   if !d.isUrgent{
       return d.priorityLevel
   }
   return 50
}

func (g groupMessage) importance() int{
    return g.priorityLevel
}

func (s systemAlert) importance() int{
    return 100
}

// ?

func processNotification(n notification) (string, int) {
	// ?
    switch notify := n.(type) {
    case directMessage:
        return notify.senderUsername, notify.importance()
    case groupMessage:
        return notify.groupName, notify.importance()
    case systemAlert:
        return notify.alertCode, notify.importance()
    default:
        return "", 0
    }
}

