<script lang="ts">
  import CircularProgress from "@smui/circular-progress";
</script>

<spinner>
  <CircularProgress style="height: 4rem; width: 4rem;" indeterminate />
  <p><slot /></p>
</spinner>

<style>
  spinner {
    display: block;
    grid-column: span 2;
    text-align: center;
    margin: auto;
    padding: 4rem 2rem;
  }
  p {
    margin-top: 1rem;
  }
</style>
