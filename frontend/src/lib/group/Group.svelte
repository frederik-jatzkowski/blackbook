<script lang="ts">
  import Form from "$lib/util/Form.svelte";
  import Textfield from "@smui/textfield";
  import client from "$lib/api/client";
  import Button, { Label, Icon } from "@smui/button";
  import Card, {
    Actions as CardActions,
    Content as CardContent,
  } from "@smui/card";
  import Dialog, {
    Content as DialogContent,
    Actions as DialogActions,
    Title,
  } from "@smui/dialog";
  import IconButton from "@smui/icon-button";

  export let group: client.group.types.GroupData;
  let invite: client.group.types.InviteData = {
    groupId: group.id,
    userEmail: "",
    message: "",
  };
  function handleInvite() {
    client.group.invite(invite);
    invite = {
      groupId: group.id,
      userEmail: "",
      message: "",
    };
    mode = "view";
  }

  let mode: "view" | "update" | "invite" | "leave" = "view";
</script>

<Card>
  <CardContent>
    <h3>{group.name}</h3>
    <p>{group.description}</p>
  </CardContent>
  <CardActions>
    <IconButton
      class="material-icons"
      title="update"
      on:click={() => (mode = "update")}
    >
      edit
    </IconButton>
    <IconButton
      class="material-icons"
      title="invite"
      on:click={() => (mode = "invite")}
    >
      add
    </IconButton>
    <IconButton
      class="material-icons"
      title="leave"
      on:click={() => (mode = "leave")}
    >
      delete
    </IconButton>
  </CardActions>
</Card>

{#if mode == "update"}
  <Dialog open on:SMUIDialog:closed={() => (mode = "view")}>
    <DialogContent>
      <Form
        name="Gruppe bearbeiten"
        btn="Speichern"
        on:submit={() => client.group.update(group)}
      >
        <p>Bitte geben Sie der Gruppe einen Namen und eine Beschreibung.</p>
        <Textfield name="name" label="Name" bind:value={group.name} />
        <Textfield
          textarea
          bind:value={group.description}
          name="description"
          label="Beschreibung"
        />
      </Form>
    </DialogContent>
  </Dialog>
{:else if mode == "invite"}
  <Dialog open on:SMUIDialog:closed={() => (mode = "view")}>
    <DialogContent>
      <Form name="Nutzer einladen" btn="Einladen" on:submit={handleInvite}>
        <p>Bitte geben Sie der Gruppe einen Namen und eine Beschreibung.</p>
        <Textfield
          name="email"
          label="Emailaddresse"
          bind:value={invite.userEmail}
        />
        <Textfield
          textarea
          bind:value={invite.message}
          name="message"
          label="Nachricht"
        />
      </Form>
    </DialogContent>
  </Dialog>
{:else if mode == "leave"}
  <Dialog open on:SMUIDialog:closed={() => (mode = "view")}>
    <Title>Bestätigen</Title>
    <DialogContent>Wollen Sie diese Gruppe wirklich verlassen?</DialogContent>
    <DialogActions>
      <Button variant="raised" on:click={() => (mode = "view")}>
        <Label>Nein</Label>
      </Button>
      <Button on:click={() => client.group.leave(group)}>
        <Label>Ja</Label>
      </Button>
    </DialogActions>
  </Dialog>
{/if}
