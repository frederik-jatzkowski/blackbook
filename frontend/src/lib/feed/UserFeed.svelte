<script lang="ts">
    import Button3 from "$lib/interact/Button3.svelte";
import Task from "./Task.svelte";
import type { TaskData, UserFeedData } from "./types";


    export let user: UserFeedData;
    let edit: TaskData | null = null;
</script>

<h2>{user.firstName} {user.lastName}</h2>
<Button3>Neue Aufgabe</Button3>
{#each user.feed as task}
    <Task {task} edit={edit==task} on:editStart={()=>edit=task}/>
{/each}

<style>
    h2 {
        grid-column: span 2;
    }
</style>