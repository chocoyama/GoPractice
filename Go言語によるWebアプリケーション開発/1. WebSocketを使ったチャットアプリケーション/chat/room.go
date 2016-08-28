package main
type room struct {
    /*
    複数のgoroutineがマップを同時に変更する可能性があり、
    メモリの破壊やその他の予期せぬ状態がもたらされるため、チャネルを利用する
    */

    // 他のクライアントに転送するためのメッセージを保持するチャネル
    forward chan []byte
    // チャットルームに参加しようとしているクライアントのためのチャネル
    join chan *client
    // チャットルームから退出しようとしているクライアントのためのチャネル
    leave chan *client
    // clientsには在室しているすべてのクライアントが保持される
    clients map[*client]bool
}

func (r *room) run() {
    for {
        select {
        case client := <-r.join:
            r.clients[client] = true // 参加
        case client := <-r.leave:
            delete(r.clients, client) // 退出
            close(client.send)
        case msg := <-r.forward:
            // 全てのクライアントにメッセージを転送
            for client := range r.clients {
                select {
                case client.send <- msg:
                    // メッセージを送信
                default:
                    delete(r.clients, client)
                    close(client.send)
                }
            }
        }
    }
}
