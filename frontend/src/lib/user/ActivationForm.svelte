<script lang="ts">
  import Form from "$lib/util/Form.svelte";
  import client from "$lib/api/client";
  import Textfield from "@smui/textfield";
  import Icon from "@smui/textfield/icon";

  const data: client.user.types.ActivateData = {
    activationCode: "",
  };
</script>

<Form
  name="Konto aktivieren"
  btn="Aktivierung abschließen"
  on:submit={() => client.user.activate(data)}
>
  <p>
    Nach ihrer Registrierung sollten Ihren Aktivierungscode per Email erhalten
    haben. Bitte geben Sie diesen hier ein, um die Registrierung abzuschließen
    und Ihre Emailadresse zu bestätigen.
  </p>
  <Textfield
    bind:value={data.activationCode}
    name="activationCode"
    label="Aktivierungscode"
  >
    <Icon class="material-icons" slot="leadingIcon">key</Icon>
  </Textfield>
</Form>
