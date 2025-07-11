import { createBrowserRouter } from 'react-router';
import PrivateRoute from 'component/common/route/private_route.jsx';

const App = (await import('src/app')).default;
const BlankLayout = (await import('component/common/layout/blank')).default;
const TenantLayout = (await import('component/common/layout/tenant')).default;
const MainLayout = (await import('component/common/layout/main')).default;

const NotMatch = (await import('component/common/route/not_match')).default;
const AuthError = (await import('component/common/result/auth_error')).default;
const VerifyEmail = (await import('component/common/result/verify_email')).default;
const Login = (await import('component/auth/login')).default;
const Signup = (await import('component/auth/signup')).default;
const TenantDetail = (await import('component/account/tenant/detail')).default;
const TenantSetting = (await import('component/account/tenant/setting')).default;
const Role = (await import('component/account/role')).default;
const User = (await import('component/account/user')).default;
const Profile = (await import('component/account/profile')).default;
const Variable = (await import('component/config/variable')).default;
const Tenant = (await import('component/account/tenant')).default;
const Workspace = (await import('component/pm/workspace')).default;
const Project = (await import('component/pm/project')).default;
const Task = (await import('component/pm/task')).default;
const Message = (await import('component/event/message')).default;
const Doc = (await import('component/document/doc')).default;


const router = createBrowserRouter([
    {
        path: '/',
        element: <App />,
        children: [
            {
                path: 'auth-error',
                lazy: async () => ({ Component: AuthError })
            },
            {
                path: 'verify-email',
                lazy: async () => ({ Component: VerifyEmail })
            },
            {
                path: 'login',
                element: <BlankLayout />,
                children: [
                    {
                        path: '',
                        lazy: async () => ({ Component: Login })
                    }
                ]
            },
            {
                path: 'signup',
                element: <BlankLayout />,
                children: [
                    {
                        path: '',
                        lazy: async () => ({ Component: Signup })
                    }
                ]
            },
            {
                path: 'account/tenant/:tenant_id',
                element: <TenantLayout />,
                children: [
                    {
                        path: '',
                        element: <PrivateRoute />,
                        children: [
                            {
                                path: '',
                                lazy: async () => ({ Component: TenantDetail })
                            },
                            {
                                path: 'role',
                                lazy: async () => ({ Component: Role })
                            },
                            {
                                path: 'user',
                                lazy: async () => ({ Component: User })
                            }
                        ]
                    }
                ]
            },
            {
                path: '',
                element: <MainLayout />,
                children: [
                    {
                        path: '',
                        element: <PrivateRoute />,
                        children: [
                            {
                                path: '',
                                lazy: async () => ({ Component: Profile })
                            },
                            {
                                path: '/account/tenant/setting',
                                lazy: async () => ({ Component: TenantSetting })
                            },
                            {
                                path: 'config/variable',
                                lazy: async () => ({ Component: Variable })
                            },
                            {
                                path: 'account/tenant',
                                lazy: async () => ({ Component: Tenant })
                            },
                            {
                                path: 'account/role',
                                lazy: async () => ({ Component: Role })
                            },
                            {
                                path: 'account/user',
                                lazy: async () => ({ Component: User })
                            },
                            {
                                path: 'pm/workspace',
                                lazy: async () => ({ Component: Workspace })
                            },
                            {
                                path: 'pm/project',
                                lazy: async () => ({ Component: Project })
                            },
                            {
                                path: 'pm/project/:projectId',
                                lazy: async () => ({ Component: Task })
                            },
                            {
                                path: 'pm/task/:taskId',
                                lazy: async () => ({ Component: Message })
                            },
                            {
                                path: 'pm/task/:taskId/doc/:docId?',
                                lazy: async () => ({ Component: Doc })
                            }
                        ]
                    }
                ]
            },
            {
                path: '*',
                lazy: async () => ({ Component: NotMatch })
            }
        ]
    }
]);
export default router;
