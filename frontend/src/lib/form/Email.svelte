<script lang="ts">
    export let name: string = "email";
    export let placeholder: string = "Emailadresse";
    export let value: string;
</script>

<input type="email" {name} {placeholder} bind:value/>

<style>
    input {
        grid-column: span 2;
        font-size: 1.2rem;
        padding: .6rem;
        border: none;
        outline: none;
        background-color: var(--COLOR-2);
        color: var(--COLOR-4);
    }
    ::placeholder { /* Chrome, Firefox, Opera, Safari 10.1+ */
        color: var(--COLOR-3);
        opacity: 1; /* Firefox */
    }
</style>