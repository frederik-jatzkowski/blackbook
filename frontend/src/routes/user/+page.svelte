<script lang="ts">
  import DeleteForm from "$lib/user/DeleteForm.svelte";
  import EditForm from "$lib/user/EditForm.svelte";
  import { session } from "$lib/api/client";
  import Button, { Label, Icon } from "@smui/button";
  import Dialog, { Content } from "@smui/dialog";

  let mode: "view" | "edit" | "delete" = "view";
</script>

<h2>Ihr Nutzerkonto</h2>

<p>{$session.payload?.firstName} {$session.payload?.lastName}</p>
<p>{$session.payload?.email}</p>

<Button variant="raised" on:click={() => (mode = "edit")}>
  <Icon class="material-icons">edit</Icon>
  <Label>Nutzerdaten ändern</Label>
</Button>

<Button variant="raised" on:click={() => (mode = "delete")}>
  <Icon class="material-icons">delete</Icon>
  <Label>Konto löschen</Label>
</Button>

<Dialog open={mode == "edit"} on:SMUIDialog:closed={() => (mode = "view")}>
  <Content>
    <EditForm />
  </Content>
</Dialog>
<Dialog open={mode == "delete"} on:SMUIDialog:closed={() => (mode = "view")}>
  <Content>
    <DeleteForm />
  </Content>
</Dialog>
