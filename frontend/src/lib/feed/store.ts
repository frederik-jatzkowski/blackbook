import { testResponse } from "$lib/test_response";
import { writable, type Writable } from "svelte/store";
import type { FeedData } from "./types";

export const feedStore: Writable<FeedData|null> = writable(testResponse)