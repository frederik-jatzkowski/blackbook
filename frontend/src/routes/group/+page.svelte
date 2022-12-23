<script lang="ts">
  import Group from "$lib/group/Group.svelte";
  import GroupCreationForm from "$lib/group/GroupCreationForm.svelte";
  import Invitation from "$lib/group/Invitation.svelte";
  import client, { groupFeed } from "$lib/api/client";
  import Messages from "$lib/util/Messages.svelte";
  import Spinner from "$lib/util/Spinner.svelte";
</script>

{#await client.group.getAll()}
  <Spinner>Gruppen werden geladen</Spinner>
{:then}
  {#if $groupFeed.payload}
    <Messages store={groupFeed} />
    {#if $groupFeed.payload.invitations}
      <h2>Einladungen</h2>
      {#each $groupFeed.payload.invitations as invitation}
        <Invitation {invitation} />
      {/each}
    {/if}

    <h2>Ihre Gruppen</h2>

    <GroupCreationForm />

    {#if $groupFeed.payload.groups}
      {#each $groupFeed.payload.groups as group}
        <Group {group} />
      {/each}
    {:else}
      <p>Du gehörst keinen Gruppen an.</p>
    {/if}
  {/if}
{/await}
