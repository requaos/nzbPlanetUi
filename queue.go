package main

import (
	"fmt"
	"time"

	"github.com/labstack/gommon/log"
)

var (
	finished = false
)

func GetQueueDetails(queue *QueueModel) {
	sabQueue, err := SABnzbd.AdvancedQueue(0, 10)
	if err != nil {
		log.Error("Error reading SABnzbd queue")
	}
	if startingUp || !finished || len(sabQueue.Slots) > 0 {
		ClearQueue(queue)
		for i := 0; i < len(sabQueue.Slots); i++ {
			var q = NewQueue(nil)
			q.SetName(sabQueue.Slots[i].Filename)
			q.SetSize(sabQueue.Slots[i].Size)
			q.SetRemaining(sabQueue.Slots[i].Percentage + "%")
			q.SetItemStatus(sabQueue.Slots[i].Status)
			queue.AddQueue(q)
		}
		sabHistory, err := SABnzbd.History(0, 90)
		if err != nil {
			log.Error("Error reading SABnzbd history")
		}
		stillWorking := false
		for b := 0; b < len(sabHistory.Slots); b++ {
			var q = NewQueue(nil)
			q.SetName(sabHistory.Slots[b].Name)
			q.SetSize(sabHistory.Slots[b].Size)
			q.SetItemStatus(sabHistory.Slots[b].Status)
			q.SetStorage(fmt.Sprintf("file://%s", sabHistory.Slots[b].Storage))
			q.SetRemaining(time.Unix(int64(sabHistory.Slots[b].Completed), 0).Format("Mon, 02 Jan 2006 15:04:05"))
			queue.AddQueue(q)
			if sabHistory.Slots[b].Status != "Completed" {
				stillWorking = true
			}
		}
		finished = !stillWorking && len(sabQueue.Slots) == 0
		startingUp = false
	}
}

func ClearQueue(queue *QueueModel) {
	queue.BeginResetModel()
	queue.SetItems([]*Queue{NewQueue(nil)})
	queue.EndResetModel()
	queue.RemoveQueue(0)
}

func LoopLoadQueue(q *QueueModel) {
	tick := time.NewTicker(time.Millisecond * 500)
	for {
		select {
		case <-tick.C:
			GetQueueDetails(q)
		}
	}
}

func Round(d, r time.Duration) time.Duration {
	if r <= 0 {
		return d
	}
	neg := d < 0
	if neg {
		d = -d
	}
	if m := d % r; m+m < r {
		d = d - m
	} else {
		d = d + r - m
	}
	if neg {
		return -d
	}
	return d
}
