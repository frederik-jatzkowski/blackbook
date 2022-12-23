<script lang="ts">
  import Form from "$lib/util/Form.svelte";
  import Textfield from "@smui/textfield";
  import Icon from "@smui/textfield/icon";
  import client, { session } from "$lib/api/client";

  const data: client.user.types.LoginData = {
    email: "",
    password: "",
  };
</script>

<Form
  name="Löschung bestätigen"
  btn="Konto löschen"
  on:submit={() => client.user.del(data)}
>
  <p>Bitte bestätigen Sie die Löschung ihres Kontos mit Ihren Zugangsdaten.</p>
  <Textfield bind:value={data.email} type="email" label="Emailaddresse">
    <Icon class="material-icons" slot="leadingIcon">mail</Icon>
  </Textfield>
  <Textfield bind:value={data.password} type="password" label="Passwort">
    <Icon class="material-icons" slot="leadingIcon">key</Icon>
  </Textfield>
</Form>
