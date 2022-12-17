<script lang="ts">
    import Button from "./Button.svelte";

    export let span : number = 2;
</script>

<Button {span} bg="var(--COLOR-2)" color="var(--COLOR-4)" on:click><slot/></Button>