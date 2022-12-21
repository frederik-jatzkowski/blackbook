<script lang="ts">
    export let span : number = 2;
    export let href : string = "/";
</script>


<a style={`grid-column: span ${span};`} {href}>
    <slot/>
</a>

<style>
    a {
        cursor: pointer;
        padding: calc(var(--SPACING) * 0.6);
        background-color: var(--COLOR-2);
        outline: none;
        border: none;
        color: inherit;
        text-transform: uppercase;
        text-decoration: none;
        text-align: center;
        font-size: 1rem;
        /* box-shadow: inset 0px 0px 0px -3px var(--COLOR-3); */
        transition: all .1s ease;
    }
    /* a:focus  {
        box-shadow: inset 0px -5px 0px -3px var(--COLOR-3);
    } */
</style>