<script lang="ts">
    import Button from "./Button.svelte";

    export let span : number = 2;
</script>

<Button {span} bg="var(--COLOR-3)" color="var(--COLOR-1)" on:click><slot/></Button>