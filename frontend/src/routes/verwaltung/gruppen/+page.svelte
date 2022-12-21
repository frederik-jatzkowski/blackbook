<script>
    import Group from "$lib/group/Group.svelte";
  import GroupCreate from "$lib/group/GroupCreate.svelte";
    import Invitation from "$lib/group/Invitation.svelte";
import Anchor from "$lib/interact/Anchor.svelte";
    import Button from "$lib/interact/Button.svelte";
    import Button3 from "$lib/interact/Button3.svelte";
    import client, { groupFeed } from "$lib/util/api/client";
  import Error from "$lib/util/Error.svelte";
  import Errors from "$lib/util/Errors.svelte";
  import Spinner from "$lib/util/Spinner.svelte";
  import Success from "$lib/util/Success.svelte";
</script>

<Anchor href="/verwaltung" span={1}>Zurück</Anchor>
<h3>Verwaltung | Gruppen</h3>

{#await client.group.getAll()}
    <Spinner>Gruppen werden geladen</Spinner>
{:then}
    {#if $groupFeed.payload}
        <Errors response={$groupFeed}/>
        {#if $groupFeed.payload.invitations}
            <h3>Einladungen</h3>
            {#each $groupFeed.payload.invitations as invitation}
                <Invitation {invitation}/>
            {/each}
        {/if}

        <h3>Meine Gruppen</h3>

        <GroupCreate/>

        {#if $groupFeed.payload.groups}
            {#each $groupFeed.payload.groups as group}
                <Group {group}/>
            {/each}
        {:else}
            <p>Du gehörst keinen Gruppen an.</p>
        {/if}
    {/if}
{/await}



