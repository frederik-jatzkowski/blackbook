<script lang="ts">
    export let name: string = "password";
    export let placeholder: string = "Passwort";
    export let value: string;
</script>

<input type="password" {name} {placeholder} bind:value/>

<style>
    input {
        grid-column: span 2;
        font-size: 1.2rem;
        padding: calc(var(--SPACING) * 0.6);
        border: solid .1rem var(--COLOR-1);
        outline: none;
        background-color: var(--COLOR-2);
        color: var(--COLOR-4);
    }
    ::placeholder { /* Chrome, Firefox, Opera, Safari 10.1+ */
        color: var(--COLOR-3);
        opacity: 1; /* Firefox */
    }
</style>