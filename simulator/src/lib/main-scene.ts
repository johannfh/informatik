import * as ex from "excalibur";
import { writable, type Readable, type Writable } from "svelte/store";
import type { ActorGetter, AnimalOption } from "./types";
import { Sheep } from "./sheep";
import { Wolf } from "./wolf";
import { Fox } from "./fox";
import { Chicken } from "./chicken";
import { AUTO_SPAWN, CHICKEN_COOLDOWN, CHICKEN_SPEED, FOX_COOLDOWN, FOX_SPEED, screenHeight, screenWidth, SHEEP_COOLDOWN, SHEEP_SPEED, WOLF_COOLDOWN, WOLF_SPEED } from "./constants";
import { defaultSettings, type Settings } from "./settings";

export type SpawnAnimalEvent = {
    type: "spawnAnimalEvent";
    animal: AnimalOption;
    position: ex.Vector;
};
export type KillAnimalEvent = {
    type: "killAnimalEvent";
    position: ex.Vector;
};

export type InputEvent = SpawnAnimalEvent | KillAnimalEvent;

export let inputEvents: Writable<InputEvent[]> = writable([]);

export class MainScene extends ex.Scene {
    private actorGetter: ActorGetter = (func) => this.actors.filter((a) => func(a));
    private foxCooldown = FOX_COOLDOWN;
    private wolfCooldown = WOLF_COOLDOWN;
    private sheepCooldown = SHEEP_COOLDOWN;
    private chickenCooldown = CHICKEN_COOLDOWN;
    private settings: Settings = defaultSettings;
    
    constructor(private globalSettings: Readable<Settings>) {
        super();
        this.globalSettings.subscribe(v => this.settings = v)
    }

    override onInitialize(engine: ex.Engine): void {
        inputEvents.subscribe((v) => {
            let event = v.shift();
            if (!event) return;
            console.log(`handling ${event.type}`);
            this.handleInputEvent(event);
        });
    }

    private handleInputEvent(event: InputEvent) {
        if (event.type === "spawnAnimalEvent") {
            this.handleSpawnAnimal(event);
            return;
        }

        console.error(`event handler for '${event.type}' currently unimplemented`);
    }

    override onPostUpdate(engine: ex.Engine, elapsedMs: number): void {
        if (this.settings.autoSpawn) this.autoSpawn(elapsedMs);
    }

    autoSpawn(elapsedMs: number) {
        this.foxCooldown -= elapsedMs;
        if (this.foxCooldown < 0) {
            this.add(
                new Fox({
                    position: ex.vec(Math.random() * screenWidth, Math.random() * screenHeight),
                    actorGetter: this.actorGetter,
                    speed: FOX_SPEED,
                }),
            );
            this.foxCooldown = FOX_COOLDOWN;
        }

        this.wolfCooldown -= elapsedMs;
        if (this.wolfCooldown < 0) {
            this.add(
                new Wolf({
                    position: ex.vec(Math.random() * screenWidth, Math.random() * screenHeight),
                    actorGetter: this.actorGetter,
                    speed: WOLF_SPEED,
                }),
            );
            this.wolfCooldown = WOLF_COOLDOWN;
        }

        this.chickenCooldown -= elapsedMs;
        if (this.chickenCooldown < 0) {
            this.add(
                new Chicken({
                    position: ex.vec(Math.random() * screenWidth, Math.random() * screenHeight),
                    actorGetter: this.actorGetter,
                    speed: CHICKEN_SPEED,
                }),
            );
            this.chickenCooldown = CHICKEN_COOLDOWN;
        }

        this.sheepCooldown -= elapsedMs;
        if (this.sheepCooldown < 0) {
            this.add(
                new Sheep({
                    position: ex.vec(Math.random() * screenWidth, Math.random() * screenHeight),
                    actorGetter: this.actorGetter,
                    speed: SHEEP_SPEED,
                }),
            );
            this.sheepCooldown = SHEEP_COOLDOWN;
        }
    }

    private handleSpawnAnimal({ animal, position }: SpawnAnimalEvent) {
        const actorGetter = this.actorGetter;

        switch (animal) {
            case Wolf:
                this.add(new Wolf({ position, speed: WOLF_SPEED, actorGetter }));
                break;

            case Sheep:
                this.add(new Sheep({ position, speed: SHEEP_SPEED, actorGetter }));
                break;

            case Fox:
                this.add(new Fox({ position, speed: FOX_SPEED, actorGetter }));
                break;

            case Chicken:
                this.add(new Chicken({ position, speed: CHICKEN_SPEED, actorGetter }));
                break;
        }
    }
}
