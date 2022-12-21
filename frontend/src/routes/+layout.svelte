<script lang="ts">
    import Anchor from "$lib/interact/Anchor.svelte";
    import ActivationForm from "$lib/user/ActivationForm.svelte";
    import CreationForm from "$lib/user/CreationForm.svelte";
    import LoginForm from "$lib/user/LoginForm.svelte";
    import client, { session } from "$lib/util/api/client"
    import Spinner from "$lib/util/Spinner.svelte";
    import Errors from "$lib/util/Errors.svelte";

    let messageHeight = 0;

</script>

<spacer style={`height: calc(${messageHeight}px - var(--SPACING) * 1.5);`}/>

{#await client.user.sessionCheck()}
    <Spinner>Blackbook wird gestartet</Spinner>
{:then}
    {#if $session.errors || $session.success}
        <messages bind:clientHeight={messageHeight}>
            <Errors response={$session}/>
        </messages>
    {/if}
    {#if $session.payload}
        {#if $session.payload.active}
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

<style>
    spacer {
        grid-column: span 2;
    }
    messages {
        position: fixed;
        display: grid;
        grid-column: span 2;
        padding: var(--SPACING);
        background-color: var(--COLOR-1);
        top: 0;
        right: 0;
        left: 0;
        height: max-content;
    }
</style>