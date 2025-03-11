permit (
    principal is CareGiver,
    action,
    resource == Capability::"UserProfile"
)
when {
    principal.cr.contains(context.crId)
    // &&
    // context.CRCGActionSet has action
    // resource in context.CRCGResourceSet
};

