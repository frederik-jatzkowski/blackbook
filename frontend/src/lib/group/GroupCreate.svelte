<script lang="ts">
    import Form from "$lib/form/Form.svelte";
  import TextArea from "$lib/form/TextArea.svelte";
    import TextField from "$lib/form/TextField.svelte";
  import Button3 from "$lib/interact/Button3.svelte";
  import client, { session } from "$lib/util/api/client";

    let create: boolean = false;

    let data = {
        name: "",
        description: ""
    }

    async function handleSubmit() {
        await client.group.create(data)
        if($session.ok) create = false;
    }
    async function startCreation() {
        data = {name: "", description: ""}
        create=true
    }
</script>

{#if create}
    <Form name="Neue Gruppe" btn="Erstellen" on:submit={handleSubmit}>
        <TextField name="name" placeholder="Name" bind:value={data.name}/>
        <TextArea bind:value={data.description} placeholder="Beschreibung" name="description"/>
    </Form>
{:else}
    <Button3 on:click={startCreation}>Neue Gruppe</Button3>
{/if}