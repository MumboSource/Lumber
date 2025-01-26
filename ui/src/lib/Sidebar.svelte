<script>
	import SidebarItem from './SidebarItem.svelte';
</script>

<div class="w-72 border-[#E0E0E0] border-r-2 flex flex-col pt-2">
    <SidebarItem url="/exe-info" iconName="appinfo" text="Application Info" />
    <SidebarItem url="/sus-exe" iconName="suspicious" text="Executables"/>
    <SidebarItem url="/activity-analysis" iconName="activity" text="Activity Analysis"/>
    <SidebarItem url="/file-events" iconName="file" text="File Events"/>
    <SidebarItem url="/sys-info" iconName="system-information" text="System Info"/>
    <SidebarItem url="/flow-chart" iconName="flowchart" text="Flow Chart"/>

    <div class="flex-grow w-full border-[#E0E0E0] border-b-2"></div>

    <SidebarItem url="/about" iconName="about" text="About"/>
</div>