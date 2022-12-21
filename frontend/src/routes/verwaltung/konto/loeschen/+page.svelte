<script lang="ts">
    import Email from "$lib/form/Email.svelte";
    import Form from "$lib/form/Form.svelte";
    import Password from "$lib/form/Password.svelte";
    import Anchor from "$lib/interact/Anchor.svelte";
    import client from "$lib/util/api/client";

    const data: client.user.types.LoginData = {
        email: "",
        password: ""
    }
</script>

<Anchor href="/verwaltung/konto" span={1}>Zurück</Anchor>
<h3>Verwaltung | Mein Konto | löschen</h3>

<Form name="Konto löschen" btn="Bestätigen" on:submit={()=>client.user.del(data)}>
    <p>Bitte geben Sie ihre Zugangsdaten ein, um die Löschung zu bestätigen.</p>
    <Email bind:value={data.email}/>
    <Password bind:value={data.password}/>
</Form>