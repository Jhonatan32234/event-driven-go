// /infrastructure/adapters/websocketAdapter.js
class WebSocketAdapter {
    constructor() {
      this.socket = null;
      this.subscribers = [];
    }
  
    connect(url) {
      this.socket = new WebSocket(url);
      console.log("Connected to WebSocket ",this.socket);
      
      this.socket.onmessage = (event) => {
        console.log("pasa aqui 2",this.socket);
        
        const data = JSON.parse(event.data);
        console.log("Received data from WebSocket", data);
        
        this.notifySubscribers(data);
      };
    }
  
    subscribe(callback) {
      this.subscribers.push(callback);
    }
  
    notifySubscribers(data) {
      this.subscribers.forEach((callback) => callback(data));
    }
  }
  
  export default WebSocketAdapter;
  