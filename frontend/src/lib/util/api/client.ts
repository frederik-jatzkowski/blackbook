import { browser } from "$app/environment"
import { env } from "$env/dynamic/public"
import { writable, type Writable } from "svelte/store"

console.log(`using api on ${env.PUBLIC_REST_ADDRESS}`)

function buildApiUrl(path:string):string {
    return `${env.PUBLIC_REST_ADDRESS}/${path}`
}

async function fetchResponse(path:string, init?: RequestInit): Promise<client.types.Response> {
    try {
        const response = await fetch(buildApiUrl(path), init)

        return await response.json()
    } catch (error) {
        console.log(error)

        return {
            ok: false,
            errors: ["Server nicht erreichbar."]
        }
    }
}

export const session: Writable<client.types.Response> = writable({
    ok: false,
    activityError: false,
    errors: []
})
namespace client {
    export namespace types {
        export type Response = {
            ok: boolean,
            errors: string[],
            user?: user.types.User
        }
    }
    export namespace user {
        export namespace types {
            export type User = {
                id: number,
                active: boolean,
                firstName: string,
                lastName: string,
                email: string
            }
            export type LoginData = {
                email: string,
                password: string
            }
            export type CreateData = {
                firstName: string,
                lastName: string,
                email: string,
                password: string
                passwordRepeat: string
            }
            export type ActivateData = {
                activationCode: string
            }
        }
        export async function create(data: types.CreateData) {
            if(!browser) return;
            session.set(await fetchResponse(
                "user/create",
                {
                    method: "POST",
                    body: JSON.stringify(data),
                    credentials: "include"
                }
            ))
        }
        export async function activate(data: types.ActivateData) {
            if(!browser) return;
            session.set(await fetchResponse(
                "user/activate",
                {
                    method: "POST",
                    body: JSON.stringify(data),
                    credentials: "include"
                }
            ))
        }
        export async function login(data: types.LoginData) {
            if(!browser) return;
            session.set(await fetchResponse(
                "user/login",
                {
                    method: "POST",
                    body: JSON.stringify(data),
                    credentials: "include"
                }
            ))
        }
        export async function logout() {
            if(!browser) return;
            session.set(await fetchResponse(
                "user/logout",
                {
                    method: "GET",
                    credentials: "include"
                }
            ))
        }
        export async function sessionCheck() {
            if(!browser) return;
            session.set(await fetchResponse(
                "user/sessionCheck",
                {
                    method: "GET",
                    credentials: "include"
                }
            ))
        }
    }
    export namespace group {
        export namespace types {
            
        }
        export async function create() {
            
        }
        export async function invite() {
            
        }
        export async function accept() {
            
        }
        export async function decline() {
            
        }
        export async function update() {
            
        }
        export async function exit() {
            
        }
    }
}

export default client