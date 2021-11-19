<script>
	import Stock from '../../artifacts/contracts/Stock.sol/Stock.json'
	import { ethers } from 'ethers'
	import { onMount } from 'svelte'

	export let nbrOwners = 0;

</script>

<div class="card">
  <div class="card-header">
    Number of stoks's owner
  </div>
  <ul class="list-group list-group-flush">
    <li class="list-group-item">{nbrOwners}</li>
  </ul>
</div>

<style>

</style>
