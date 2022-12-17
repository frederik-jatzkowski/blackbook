<script>
    import Anchor from "$lib/interact/Anchor.svelte";
    import Button3 from "$lib/interact/Button3.svelte";
    import client from "$lib/util/api/client";
</script>

<Anchor href="/verwaltung" span={1}>Zurück</Anchor>
<Button3 on:click={()=>client.user.logout()}>Ausloggen</Button3>