import { writable } from "svelte/store";

export const defaultSettings: Settings = {
    autoSpawn: false
}

export type Settings = {
    autoSpawn: boolean;
}

export let settings = writable<Settings>(defaultSettings)