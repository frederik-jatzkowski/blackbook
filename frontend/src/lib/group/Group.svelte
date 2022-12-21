<script lang="ts">
    import Form from "$lib/form/Form.svelte";
    import TextArea from "$lib/form/TextArea.svelte";
    import TextField from "$lib/form/TextField.svelte";
    import Button3 from "$lib/interact/Button3.svelte";
    import client from "$lib/util/api/client";


    export let group: client.group.types.GroupData;
    let invite: client.group.types.InviteData = {
        groupId: group.id,
        userEmail: "",
        message: "",
    }
    function handleInvite() {
        client.group.invite(invite);
        invite = {
            groupId: group.id,
            userEmail: "",
            message: "",
        }
        mode="view";
    }

    let mode: "view" | "update" | "invite" | "leave" = "view"
</script>

<group>
    {#if mode == "view"}
        <h3>{group.name}</h3>
        <p>{group.description}</p>
        <Button3 span={1} on:click={()=>mode = "update"}>Bearbeiten</Button3>
        <Button3 span={1} on:click={()=>mode = "invite"}>Einladen</Button3>
        <Button3 span={1} on:click={()=>mode = "leave"}>Verlassen</Button3>
    {:else if mode == "update"}
        <Form name="Gruppe bearbeiten" btn="Speichern" on:submit={()=>{client.group.update(group), mode="view"}}>
            <TextField name="name" placeholder="Name" bind:value={group.name}/>
            <TextArea name="name" placeholder="Name" bind:value={group.description}/>
        </Form>
        <Button3 on:click={()=>mode="view"}>Zurück</Button3>
    {:else if mode == "invite"}
        <Form name="Nutzer einladen" btn="Einladen" on:submit={handleInvite}>
            <TextField name="email" placeholder="Emailaddresse" bind:value={invite.userEmail}/>
            <TextField name="message" placeholder="Nachricht" bind:value={invite.message}/>
        </Form>
        <Button3 on:click={()=>mode="view"}>Zurück</Button3>
    {:else}
        <h3>{group.name}</h3>
        <p>{group.description}</p>
        <Button3 on:click={()=>{client.group.leave(group), mode="view"}}>Gruppe wirklich verlassen!</Button3>
        <Button3 on:click={()=>mode="view"}>Zurück</Button3>
    {/if}
</group>

<style>
    group {
        display: grid;
        grid-column: span 2;
        background-color: var(--COLOR-2);
        padding: var(--SPACING);
        grid-gap: var(--SPACING);
    }
</style>