<script>
	import SidebarItem from './SidebarItem.svelte';

    export let selectedIndex = 0;
</script>

<div class="w-72 border-[#E0E0E0] border-r-2 flex flex-col pt-2">
    <SidebarItem to="/application-info" iconName="appinfo" selected={selectedIndex == 0} text="Application Info"/>
    <SidebarItem to="/registry-integrity" iconName="registry" selected={selectedIndex == 1} text="Registry Integrity"/>
    <SidebarItem to="/suspicious-executables" iconName="suspicious" selected={selectedIndex == 2} text="Suspicious Executables"/>
    <SidebarItem to="/activity-analysis" iconName="activity" selected={selectedIndex == 3} text="Activity Analysis"/>
    <SidebarItem to="/file-actions" iconName="file" selected={selectedIndex == 4} text="File Actions"/>
    <SidebarItem to="/system-information" iconName="system-information" selected={selectedIndex == 5} text="System Information"/>
    <SidebarItem to="/flow-chart" iconName="flowchart" selected={selectedIndex == 6} text="Flow Chart"/>
</div>