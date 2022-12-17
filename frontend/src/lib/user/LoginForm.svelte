<script lang="ts">
    import Email from "$lib/form/Email.svelte";
    import Form from "$lib/form/Form.svelte";
    import Password from "$lib/form/Password.svelte";
    import client from "$lib/util/api/client";
    import Paragraph from "$lib/util/Paragraph.svelte";

    const data = {
        email: "",
        password: ""
    }

</script>

<Form name="Login" btn="Anmelden" on:submit={()=>client.user.login(data)}>
    <Paragraph>Bitte melden Sie sich hier an, wenn Sie bereits ein Konto besitzen. Andernfalls können Sie sich unten registrieren.</Paragraph>
    <Email bind:value={data.email}/>
    <Password bind:value={data.password}/>
</Form>