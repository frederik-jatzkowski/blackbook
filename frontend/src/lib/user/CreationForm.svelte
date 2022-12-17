<script lang="ts">
    import Email from "$lib/form/Email.svelte";
    import Form from "$lib/form/Form.svelte";
    import Password from "$lib/form/Password.svelte";
    import TextField from "$lib/form/TextField.svelte";
    import Button3 from "$lib/interact/Button3.svelte";
    import client from "$lib/util/api/client";
    import Paragraph from "$lib/util/Paragraph.svelte";

    const data: client.user.types.CreateData = {
        firstName: "",
        lastName: "",
        email: "",
        password: "",
        passwordRepeat: ""
    }

</script>


<Form name="Neu Registrieren" btn="Registrieren" on:submit={()=>client.user.create(data)}>
    <Paragraph>Hier können Sie ein neues Konto erstellen.</Paragraph>
    <TextField bind:value={data.firstName} name="firstName" placeholder="Vorname"/>
    <TextField bind:value={data.lastName} name="lastName" placeholder="Nachname"/>
    <Email bind:value={data.email}/>
    <Password bind:value={data.password}/>
    <Password bind:value={data.passwordRepeat} name="passwordRepeat" placeholder="Passwort wiederholen"/>
</Form>