import { z } from "zod";

export const MessageTypeSchema = z.enum([
    "client.game.water.add",
    "client.game.entities.add",

    "server.game.water.updated",
    "server.game.entities.updated",
    "server.game.event",

    "error.invalidMessageFormat",
    "error.unknownMessageType",
    "error.unknownEntity",
]);

export const MessageSchema = z.object({
    messageType: MessageTypeSchema,
});

const EntityTypeSchema = z.enum(["wolf", "fox", "chicken", "sheep"]);
const EntityPositionSchema = z.object({
    x: z.number(),
    y: z.number(),
});

export const EntitySchema = z.object({
    id: z.number(),
    type: EntityTypeSchema,
    position: EntityPositionSchema,
});

export type EntityType = z.infer<typeof EntityTypeSchema>;
export type EntityPosition = z.infer<typeof EntityPositionSchema>;
export type Entity = z.infer<typeof EntitySchema>;

// `client` Scope Messages

export const ClientGameEntitiesAddSchema = MessageSchema.extend({
    messageType: z.literal("client.game.entities.add"),
    entity: z.string(),
});
export const ClientGameWaterAddSchema = MessageSchema.extend({
    messageType: z.literal("client.game.water.add"),
    water: z.number(),
});

// `server` Scope Messages

export const ServerGameWaterUpdatedSchema = MessageSchema.extend({
    messageType: z.literal("server.game.water.updated"),
    water: z.number(),
});
export const ServerGameEntitiesUpdatedSchema = MessageSchema.extend({
    messageType: z.literal("server.game.entities.updated"),
    entities: z.array(EntitySchema),
});
export const ServerGameEventSchema = MessageSchema.extend({
    messageType: z.literal("server.game.event"),
    event: z.string(),
});

// `error` Scope Messages

export const ErrorInvalidMessageTypeSchema = MessageSchema.extend({
    messageType: z.literal("error.unknownMessageType"),
    message: z.string(),
});
export const ErrorInvalidMessageFormatSchema = MessageSchema.extend({
    messageType: z.literal("error.invalidMessageFormat"),
    message: z.string(),
});
export const ErrorUnknownEntitySchema = MessageSchema.extend({
    messageType: z.literal("error.unknownEntity"),
    message: z.string(),
});

export const DiscriminatedMessageSchema = z.discriminatedUnion("messageType", [
    // `client` Scope
    ClientGameWaterAddSchema,
    ClientGameEntitiesAddSchema,

    // `server` Scope
    ServerGameWaterUpdatedSchema,
    ServerGameEntitiesUpdatedSchema,
    ServerGameEventSchema,

    // `error` Scope
    ErrorInvalidMessageTypeSchema,
    ErrorInvalidMessageFormatSchema,
    ErrorUnknownEntitySchema,
]);
