import { browser } from "$app/environment"
import { env } from "$env/dynamic/public"
import type { GroupData } from "$lib/feed/types"
import { writable, type Writable } from "svelte/store"

console.log(`using api on ${env.PUBLIC_REST_ADDRESS}`)

function buildApiUrl(path:string):string {
    return `${env.PUBLIC_REST_ADDRESS}/${path}`
}

async function fetchResponse<T>(path:string, init?: RequestInit): Promise<client.types.Response<T>> {
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

export const session: Writable<client.types.Response<client.user.types.User>> = writable({
    ok: false,
    activityError: false,
    errors: []
})
export const groupFeed: Writable<client.types.Response<client.group.types.GroupFeed>> = writable({
    ok: false,
    activityError: false,
    errors: [],
    payload: {groups: [], invitations: []}
})
namespace client {
    export namespace types {
        export type Response<T> = {
            ok: boolean,
            success?: string
            errors: string[],
            payload?: T
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
            export type UpdateData = {
                firstName: string,
                lastName: string,
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
        export async function update(data: types.UpdateData) {
            if(!browser) return;
            session.set(await fetchResponse(
                "user/update",
                {
                    method: "POST",
                    body: JSON.stringify(data),
                    credentials: "include"
                }
            ))
        }
        export async function del(data: types.LoginData) {
            if(!browser) return;
            session.set(await fetchResponse(
                "user/delete",
                {
                    method: "POST",
                    body: JSON.stringify(data),
                    credentials: "include"
                }
            ))
        }
    }
    export namespace group {
        export namespace types {
            export type CreateData = {
                name: string,
                description: string
            }
            export type InviteData = {
                groupId: number,
                userEmail: string,
                message: string
            }
            export type AcceptDeclineData = {
                invitationId: number
            }
            export type InvitationData = {
                id: number,
                message: string,
                senderFirstName: string,
                senderEmail: string,
                groupName: string,
                groupDescription: string
            }
            export type GroupData = {
                id: number,
                name: string,
                description: string
            }
            export type GroupFeed = {
                invitations: InvitationData[]
                groups: GroupData[]
            }
        }
        export async function getAll() {
            if(!browser) return;
            groupFeed.set(await fetchResponse<client.group.types.GroupFeed>(
                "group/getAll",
                {
                    method: "GET",
                    credentials: "include"
                }
            ))
        }
        export async function create(data: types.CreateData) {
            if(!browser) return;
            groupFeed.set(await fetchResponse(
                "group/create",
                {
                    method: "POST",
                    body: JSON.stringify(data),
                    credentials: "include"
                }
            ))
        }
        export async function invite(data: types.InviteData) {
            if(!browser) return;
            groupFeed.set(await fetchResponse(
                "group/invite",
                {
                    method: "POST",
                    body: JSON.stringify(data),
                    credentials: "include"
                }
            ))
        }
        export async function accept(data: types.AcceptDeclineData) {
            if(!browser) return;
            groupFeed.set(await fetchResponse(
                "group/accept",
                {
                    method: "POST",
                    body: JSON.stringify(data),
                    credentials: "include"
                }
            ))
        }
        export async function decline(data: types.AcceptDeclineData) {
            if(!browser) return;
            groupFeed.set(await fetchResponse(
                "group/decline",
                {
                    method: "POST",
                    body: JSON.stringify(data),
                    credentials: "include"
                }
            ))
        }
        export async function update(data: GroupData) {
            if(!browser) return;
            groupFeed.set(await fetchResponse(
                "group/update",
                {
                    method: "POST",
                    body: JSON.stringify(data),
                    credentials: "include"
                }
            ))
        }
        export async function leave(data: GroupData) {
            if(!browser) return;
            groupFeed.set(await fetchResponse(
                "group/leave",
                {
                    method: "POST",
                    body: JSON.stringify(data),
                    credentials: "include"
                }
            ))
        }
    }
}

export default client