import { writable } from "svelte/store";
import type { User } from "@encedeus/js-api";

export const accessTokenStore = writable("");
export const userStore = writable<User | undefined>();