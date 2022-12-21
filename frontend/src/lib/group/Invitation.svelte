<script lang="ts">
    import Button3 from "$lib/interact/Button3.svelte";
import client from "$lib/util/api/client";


    export let invitation: client.group.types.InvitationData;
</script>

<invitation>
    <p>{invitation.senderFirstName} ({invitation.senderEmail}) läd dich ein, der Gruppe {invitation.groupName} beizutreten:</p>
    <p>{invitation.message}</p>
    <p>{invitation.groupDescription}</p>
    <Button3 span={1} on:click={()=>client.group.accept({invitationId:invitation.id})}>Annehmen</Button3>
    <Button3 span={1} on:click={()=>client.group.decline({invitationId:invitation.id})}>Ablehnen</Button3>
</invitation>

<style>
    invitation {
        display: grid;
        grid-column: span 2;
        background-color: var(--COLOR-2);
        padding: var(--SPACING);
        grid-gap: var(--SPACING);
    }
</style>