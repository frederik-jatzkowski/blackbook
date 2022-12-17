<spinner>
    <span class="spinner" />
    <p><slot /></p>
</spinner>
  
  <style>
    spinner {
      grid-column: span 2;
      margin: auto;
      padding: 4rem 2rem;
    }
    .spinner {
      width: 48px;
      height: 48px;
      border-radius: 50%;
      display: block;
      margin: auto;
      border-top: 3px solid #ddd;
      border-right: 3px solid transparent;
      box-sizing: border-box;
      animation: rotation 1s linear infinite;
    }
    p {
      display: block;
      text-align: center;
      padding-top: 1rem;
    }
    @keyframes rotation {
      0% {
        transform: rotate(0deg);
      }
      100% {
        transform: rotate(360deg);
      }
    }
  </style>