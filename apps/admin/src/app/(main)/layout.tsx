import { SidebarProvider } from "@/components/ui/sidebar";
import { AppSidebar } from "@/components/Sidebar/app-sidebar";
import { Navbar } from "@/components";

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <SidebarProvider>
      <main className="flex flex-col w-full">
        <div className="flex">
          <AppSidebar />

          <div className="flex flex-col w-full">
            <Navbar />
            <div className="w-full h-max overflow-auto overflow p-2 pb-5">
              {children}
            </div>
          </div>
        </div>
      </main>
    </SidebarProvider>
  );
}
