<script lang="ts">
  import Form from "$lib/util/Form.svelte";
  import client from "$lib/api/client";
  import Textfield from "@smui/textfield";
  import Icon from "@smui/textfield/icon";

  const data: client.user.types.LoginData = {
    email: "",
    password: "",
  };
</script>

<Form name="Login" btn="Anmelden" on:submit={() => client.user.login(data)}>
  <p>
    Bitte melden Sie sich hier an, wenn Sie bereits ein Konto besitzen.
    Andernfalls können Sie sich unten registrieren.
  </p>
  <Textfield bind:value={data.email} type="email" label="Emailaddresse">
    <Icon class="material-icons" slot="leadingIcon">mail</Icon>
  </Textfield>
  <Textfield bind:value={data.password} type="password" label="Passwort">
    <Icon class="material-icons" slot="leadingIcon">key</Icon>
  </Textfield>
</Form>
