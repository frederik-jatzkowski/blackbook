<script lang="ts">
  import Form from "$lib/util/Form.svelte";
  import Textfield from "@smui/textfield";
  import Icon from "@smui/textfield/icon";
  import client, { session } from "$lib/api/client";

  const data: client.user.types.UpdateData = {
    firstName: $session.payload?.firstName ?? "",
    lastName: $session.payload?.lastName ?? "",
  };
</script>

<Form
  name="Kontodaten ändern"
  btn="Speichern"
  on:submit={() => client.user.update(data)}
>
  <p>Hier können Sie ihre Kontodaten ändern.</p>
  <Textfield bind:value={data.firstName} name="firstName" label="Vorname">
    <Icon class="material-icons" slot="leadingIcon">person</Icon>
  </Textfield>
  <Textfield bind:value={data.lastName} name="lastName" label="Nachname">
    <Icon class="material-icons" slot="leadingIcon">person</Icon>
  </Textfield>
</Form>
