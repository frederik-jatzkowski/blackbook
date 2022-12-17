<script lang="ts">
    import Form from "$lib/form/Form.svelte";
    import TextField from "$lib/form/TextField.svelte";
    import client from "$lib/util/api/client";
    import Paragraph from "$lib/util/Paragraph.svelte";

    const data: client.user.types.ActivateData = {
        activationCode: ""
    }
</script>

<Form name="Konto aktivieren" btn="Aktivierung abschließen" on:submit={()=>client.user.activate(data)}>
    <Paragraph>Nach ihrer Registrierung sollten Ihren Aktivierungscode per Email erhalten haben. Bitte geben Sie diesen hier ein, um die Registrierung abzuschließen und Ihre Emailadresse zu bestätigen.</Paragraph>
    <TextField bind:value={data.activationCode} name="activationCode" placeholder="Aktivierungscode"/>
</Form>