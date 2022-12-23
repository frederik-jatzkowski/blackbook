<script lang="ts">
  import Form from "$lib/util/Form.svelte";
  import Textfield from "@smui/textfield";
  import Button, { Label } from "@smui/button";
  import client, { session } from "$lib/api/client";
  import Dialog, { Content } from "@smui/dialog";

  let create: boolean = false;

  let data = {
    name: "",
    description: "",
  };

  async function handleSubmit() {
    await client.group.create(data);
    if ($session.ok) create = false;
  }
  async function startCreation() {
    data = { name: "", description: "" };
    create = true;
  }
</script>

{#if create}
  <Dialog open={create} on:SMUIDialog:closed={() => (create = false)}>
    <Content>
      <Form name="Neue Gruppe" btn="Erstellen" on:submit={handleSubmit}>
        <p>Bitte geben Sie der Gruppe einen Namen und eine Beschreibung.</p>
        <Textfield name="name" label="Name" bind:value={data.name} />
        <Textfield
          textarea
          bind:value={data.description}
          name="description"
          label="Beschreibung"
        />
      </Form>
    </Content>
  </Dialog>
{/if}
<Button variant="raised" on:click={startCreation}>Neue Gruppe</Button>
