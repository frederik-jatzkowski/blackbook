<script lang="ts">
  import ActivationForm from "$lib/user/ActivationForm.svelte";
  import CreationForm from "$lib/user/CreationForm.svelte";
  import LoginForm from "$lib/user/LoginForm.svelte";
  import client, { session } from "$lib/api/client";
  import Spinner from "$lib/util/Spinner.svelte";
  import TopAppBar, { Row, Section, Title } from "@smui/top-app-bar";
  import Messages from "$lib/util/Messages.svelte";
  import Snackbar, { Label, Actions } from "@smui/snackbar";
  import IconButton from "@smui/icon-button";
  import Paper from "@smui/paper";
</script>

<svelte:head>
  <title>Blackbook</title>
</svelte:head>

<Snackbar open>
  <Label>response.success</Label>
  <Actions>
    <IconButton class="material-icons" title="Dismiss">close</IconButton>
  </Actions>
</Snackbar>

<Messages store={session} />
{#await client.user.sessionCheck()}
  <Spinner>Blackbook wird gestartet</Spinner>
{:then}
  {#if $session.payload}
    {#if $session.payload.active}
      <TopAppBar variant="static" color="secondary">
        <Row>
          <Section>
            <IconButton class="material-icons" href="/">assignment</IconButton>
            <IconButton class="material-icons" href="/user">person</IconButton>
            <IconButton class="material-icons" href="/group">group</IconButton>
          </Section>
          <Section align="end">
            <IconButton class="material-icons" on:click={client.user.logout}>
              logout
            </IconButton>
          </Section>
        </Row>
      </TopAppBar>
      <main>
        <slot />
      </main>
    {:else}
      <main>
        <Paper color="secondary">
          <ActivationForm />
        </Paper>
      </main>
    {/if}
  {:else}
    <main>
      <Paper color="secondary">
        <LoginForm />
      </Paper>
      <Paper color="secondary">
        <CreationForm />
      </Paper>
    </main>
  {/if}
{/await}

<style>
  main {
    display: flex;
    flex-direction: column;
    gap: 1rem;
    padding: var(--SPACING);
  }
</style>
