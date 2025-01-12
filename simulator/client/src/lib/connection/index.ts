import { z } from "zod";
import { DiscriminatedMessageSchema, MessageTypeSchema } from "./schemas/messages";
import { WEBSOCKET_STATUS_CODES, type WebSocketStatusCode } from "$lib/utils";
export { DiscriminatedMessageSchema, MessageTypeSchema };

export type MessageType = z.infer<typeof MessageTypeSchema>;
export type Message = z.infer<typeof DiscriminatedMessageSchema>;

type MessageListener = (message: Message) => any | void;

class Client {
    private websocket: WebSocket | null = null;
    private listeners: MessageListener[] = [];
    private connected: boolean = false;

    /** Initial reconnect delay in ms */
    private reconnectDelay: number = 1000;

    private reconnectAttemts: number = 0;
    private readonly maxReconnectAttempts: number = 10;

    private readonly expectedCloseCodes: WebSocketStatusCode[] = [
        WEBSOCKET_STATUS_CODES.CLOSE_NORMAL,
    ];

    constructor(private serverURL: string | URL) {}

    public changeServer(serverURL: string | URL) {
        this.serverURL = serverURL;
        // switch connection if currently connected
        if (this.connected) {
            this.disconnect(WEBSOCKET_STATUS_CODES.CLOSE_NORMAL);
            this.connect();
        }
    }

    public connect() {
        this.websocket = new WebSocket(this.serverURL);

        this.websocket.onopen = () => this.handleOpen();
        this.websocket.onclose = (e) => this.handleClose(e);
        this.websocket.onerror = () => this.handleError();
        this.websocket.onmessage = (e) => this.handleMessage(e);

        this.connected = true;
    }

    public disconnect(code?: number, reason?: string) {
        if (!code) {
            console.warn("no close code provided");
        }
        if (!this.websocket) return;
        this.websocket.close(code, reason);
        this.websocket = null;
        this.connected = false;
    }

    public onMessage(listener: MessageListener) {
        this.listeners.push(listener);
    }

    public isConnected = () => this.connected;

    public send(message: Message) {
        let data = JSON.stringify(message);
        this.websocket?.send(data);
    }

    private handleOpen() {
        console.info("WebSocket connected");
        this.connected = true;
        this.reconnectAttemts = 0;
    }

    private handleError() {}

    private handleClose(event: CloseEvent) {
        console.warn(`WebSocket closed: ${event.reason}`);
        this.connected = false;
        // reconnect on unexpected closes
        if (!this.expectedCloseCodes.includes(event.code as WebSocketStatusCode)) {
            this.reconnect();
        }
    }

    private handleMessage(event: MessageEvent) {
        try {
            let message = DiscriminatedMessageSchema.parse(JSON.parse(event.data));
            console.info(`Received message: ${message.messageType}`);
            for (let i = 0; i < this.listeners.length; i++) {
                this.listeners[i](message);
            }
        } catch (e) {
            console.error("failed to parse message:", e);
        }
    }

    private reconnect() {
        if (this.reconnectAttemts <= this.maxReconnectAttempts) {
            this.reconnectAttemts++;
            // exponential backoff
            const delay = this.reconnectDelay * Math.pow(2, this.reconnectAttemts);
            console.info(`Reconnecting in: ${delay}ms`);
            setTimeout(() => {
                this.connect();
            }, delay);
        } else {
            console.error(`Max reconnect attempts reached: ${this.maxReconnectAttempts}`);
        }
    }
}

export default Client;
