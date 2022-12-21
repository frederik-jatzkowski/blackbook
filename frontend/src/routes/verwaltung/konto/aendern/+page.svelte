<script lang="ts">
    import Form from "$lib/form/Form.svelte";
    import TextField from "$lib/form/TextField.svelte";
    import Anchor from "$lib/interact/Anchor.svelte";
    import client, { session } from "$lib/util/api/client";


    const data: client.user.types.UpdateData = {
        firstName: $session.payload?.firstName ?? "",
        lastName:  $session.payload?.lastName ?? ""
    }
</script>

<Anchor href="/verwaltung/konto" span={1}>Zurück</Anchor>
<h3>Verwaltung | Mein Konto | ändern</h3>

<Form name="Kontodaten ändern" btn="Bestätigen" on:submit={()=>client.user.update(data)}>
    <p>Bitte geben Sie ihre Änderungen ein.</p>
    <TextField bind:value={data.firstName} name="firstName" placeholder="Vorname"/>
    <TextField bind:value={data.lastName} name="lastName" placeholder="Nachname"/>
</Form>