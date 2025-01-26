<script>
    import testIcon from '../assets/test-img.png'
    import backArrow from '../assets/left-arrow.svg'
    import { navigate } from 'svelte-routing';
</script>

<div class="flex py-3 px-5 h-20 w-full gap-5 items-center" style="border-bottom: 2px solid #E1E1E1;">
    <button on:click={() => navigate("/exe-list")} >
        <img src={backArrow} alt="select" class="nextButton w-12 rounded-[5px] transition-all" />
    </button>
</div>

<style>
    .nextButton:hover {
        background-color: #d9d9d9;
    }
</style>