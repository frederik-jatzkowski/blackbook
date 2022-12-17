<script lang="ts">
    import Button2 from "$lib/interact/Button2.svelte";
    import UserFeed from "./UserFeed.svelte";
    import type { FeedData,  UserFeedData } from "./types";
    import { feedStore } from "./store";
    import Spinner from "$lib/util/Spinner.svelte";
    
    let feed : FeedData | null = $feedStore;
    let index: number = 0;
    
    function move(change:number):number {
        if(!feed) return 0;
        return (index + feed.users.length + change) % feed.users.length
    }

    let user: UserFeedData, next: UserFeedData, prev: UserFeedData
    $: {
        if(feed) {
            user = feed.users[index]
            next = feed.users[move(1)]
            prev = feed.users[move(-1)]
        }
    }

</script>

{#if feed === null}
    <Spinner>Aufgaben werden geladen</Spinner>
{:else}
    {#if feed.users.length > 2}
        <Button2 span={1} on:click={()=>index = move(-1)}>{prev.firstName}</Button2>
        <Button2 span={1} on:click={()=>index = move(1)}>{next.firstName}</Button2>
    {:else}
        <Button2 on:click={()=>index = move(1)}>{next.firstName}</Button2>
    {/if}
    <UserFeed {user}/>
{/if}

