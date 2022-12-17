<script lang="ts">
    import Anchor from "$lib/interact/Anchor.svelte";
    import ActivationForm from "$lib/user/ActivationForm.svelte";
    import CreationForm from "$lib/user/CreationForm.svelte";
    import LoginForm from "$lib/user/LoginForm.svelte";
    import client, { session } from "$lib/util/api/client"
    import Spinner from "$lib/util/Spinner.svelte";
    import Error from "$lib/util/Error.svelte";
</script>

{#await client.user.sessionCheck()}
    <Spinner>Blackbook wird gestartet</Spinner>
{:then}
    {#if $session.errors != null}
        {#each $session.errors as err}
            <Error>{err}</Error>
        {/each}
    {/if}
    {#if $session.user}
        {#if $session.user.active}
            <Anchor href="/" span={1}>Home</Anchor>
            <slot/>
        {:else}
            <ActivationForm/>
        {/if}
    {:else}
        <LoginForm/>
        <CreationForm/>
    {/if}
{/await}
<!-- {JSON.stringify($session)} -->