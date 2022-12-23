<script lang="ts">
  import client from "$lib/api/client";

  import Card, {
    Actions as CardActions,
    Content as CardContent,
  } from "@smui/card";
  import IconButton from "@smui/icon-button";

  export let invitation: client.group.types.InvitationData;
</script>

<Card>
  <CardContent>
    <p>
      <strong>Von: </strong>
      {invitation.senderFirstName} ({invitation.senderEmail})
    </p>
    <p><strong>Gruppe: </strong>{invitation.groupName}</p>
    <p><strong>Nachricht: </strong></p>
    <p>{invitation.message}</p>
    <p><strong>Gruppenbeschreibung: </strong></p>
    <p>{invitation.groupDescription}</p>
  </CardContent>
  <CardActions>
    <IconButton
      class="material-icons"
      title="accept"
      on:click={() => client.group.accept(invitation)}
    >
      check
    </IconButton>
    <IconButton
      class="material-icons"
      title="decline"
      on:click={() => client.group.decline(invitation)}
    >
      clear
    </IconButton>
  </CardActions>
</Card>

<style>
  p {
    margin-block-end: 0.8rem;
  }
</style>
