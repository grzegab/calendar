package ws

import (
	"sync"
)

var mu sync.RWMutex

func Handle(h *Hub, msg []byte) {
	mu.Lock()
	defer mu.Unlock()

	//var data WsMsg
	//if err := json.Unmarshal(msg, &data); err != nil {
	//	fmt.Printf("failed to unmarshal: %v\n", err)
	//	return
	//}
	//
	//switch data.Type {
	//case "add_todo":
	//	id := len(todos) + 1
	//	newTodo := TodoItem{ID: id, Text: data.Text, Done: false}
	//	todos = append(todos, newTodo)
	//
	//	broadcast, _ := json.Marshal(map[string]any{
	//		"type": "todo_added",
	//		"id":   newTodo.ID,
	//		"text": newTodo.Text,
	//		"done": newTodo.Done,
	//	})
	//	h.Broadcast(broadcast)
	//
	//case "toggle_todo":
	//	for i := range todos {
	//		if todos[i].ID == data.ID {
	//			todos[i].Done = data.Done
	//			break
	//		}
	//	}
	//	broadcast, _ := json.Marshal(map[string]any{
	//		"type": "todo_toggled",
	//		"id":   data.ID,
	//		"done": data.Done,
	//	})
	//	h.Broadcast(broadcast)
	//
	//case "chat":
	//	broadcast, _ := json.Marshal(map[string]any{
	//		"type":    "chat",
	//		"message": data.Message,
	//	})
	//	h.Broadcast(broadcast)
	//}
}
