<script lang="ts">
  import Form from "$lib/util/Form.svelte";
  import Textfield from "@smui/textfield";
  import Icon from "@smui/textfield/icon";
  import client from "$lib/api/client";

  const data: client.user.types.CreateData = {
    firstName: "",
    lastName: "",
    email: "",
    password: "",
    passwordRepeat: "",
  };
</script>

<Form
  name="Neu Registrieren"
  btn="Registrieren"
  on:submit={() => client.user.create(data)}
>
  <p>Hier können Sie ein neues Konto erstellen.</p>
  <Textfield bind:value={data.firstName} name="firstName" label="Vorname">
    <Icon class="material-icons" slot="leadingIcon">person</Icon>
  </Textfield>
  <Textfield bind:value={data.lastName} name="lastName" label="Nachname">
    <Icon class="material-icons" slot="leadingIcon">person</Icon>
  </Textfield>
  <Textfield
    bind:value={data.email}
    type="email"
    name="email"
    label="Emailaddresse"
  >
    <Icon class="material-icons" slot="leadingIcon">mail</Icon>
  </Textfield>
  <Textfield
    bind:value={data.password}
    type="password"
    name="password"
    label="Passwort"
  >
    <Icon class="material-icons" slot="leadingIcon">key</Icon>
  </Textfield>
  <Textfield
    bind:value={data.passwordRepeat}
    type="password"
    name="passwordConfirm"
    label="Passwort bestätigen"
  >
    <Icon class="material-icons" slot="leadingIcon">key</Icon>
  </Textfield>
</Form>
