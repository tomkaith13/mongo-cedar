@policy_desc("This policy is used whenever a CG impersonates as a CR and accesses resources")
permit (
    principal is CareGiver,
    action,
    resource
)
when {
    context.impersonation == true &&
    principal.status == "ACCEPTED" &&
    principal.cr.contains(context.crId) &&
    context.CRCGActionSet.contains(action.name) &&
    context.CRCGResourceSet.contains(resource.name)
};

@policy_desc("This Policy is for CareGivers accessing their own resources while not impersonating.")
permit (
    principal is CareGiver, 
    action, 
    resource
)
when {
    context has impersonation &&
    context.impersonation == false &&
    principal has OwnResourceSet &&
    principal.OwnResourceSet.contains(resource.name)
};