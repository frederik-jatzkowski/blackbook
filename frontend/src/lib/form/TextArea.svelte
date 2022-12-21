<script lang="ts">

    export let name: string;
    export let placeholder: string;
    export let value: string;

</script>

<textarea bind:value rows="5" {placeholder} {name}/>

<style>
    textarea {
        grid-column: span 2;
        background-color: var(--COLOR-2);
        padding: calc(var(--SPACING) * 0.6);
        border: solid .1rem var(--COLOR-1);
        outline: none;
        resize: none;
        color: var(--COLOR-4);
        font-family: inherit;
        font-size: 1.2rem;
    }
    ::placeholder {
        color: var(--COLOR-3);
    }
</style>